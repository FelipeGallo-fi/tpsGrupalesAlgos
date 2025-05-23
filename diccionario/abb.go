package diccionario

import (
	TDAPila "tdas/pila"
)

type ab[K comparable, V any] struct {
	raiz     *nodoAb[K, V]
	cant     int
	comparar func(K, K) int
}

type nodoAb[K comparable, V any] struct {
	izq   *nodoAb[K, V]
	der   *nodoAb[K, V]
	clave K
	dato  V
}

type iteradorExternoABB[K comparable, V any] struct {
	pila TDAPila.Pila[*nodoAb[K, V]]
}

type iteradorRangoABB[K comparable, V any] struct {
	pila     TDAPila.Pila[*nodoAb[K, V]]
	desde    *K
	hasta    *K
	comparar func(K, K) int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &ab[K, V]{raiz: nil, cant: 0, comparar: funcion_cmp}
}

func (a *ab[K, V]) Pertenece(clave K) bool {
	return perteneceRecursiva(a.raiz, clave, a.comparar)
}

func perteneceRecursiva[K comparable, V any](nodo *nodoAb[K, V], clave K, comparar func(K, K) int) bool {
	if nodo == nil {
		return false
	}
	cmp := comparar(clave, nodo.clave)
	if cmp == 0 {
		return true
	} else if cmp < 0 {
		return perteneceRecursiva(nodo.izq, clave, comparar)
	} else {
		return perteneceRecursiva(nodo.der, clave, comparar)
	}
}

func (a *ab[K, V]) Cantidad() int {
	return a.cant
}

func obtenerMinimo[K comparable, V any](nodo *nodoAb[K, V]) *nodoAb[K, V] {
	if nodo == nil {
		return nil
	}
	actual := nodo
	for actual.izq != nil {
		actual = actual.izq
	}
	return actual
}

func (a *ab[K, V]) Borrar(clave K) V {
	var borrado V
	var borradoOk bool
	a.raiz, borrado, borradoOk = borrarRecursiva(a.raiz, clave, a.comparar, &a.cant)
	if !borradoOk {
		panic("La clave no pertenece al abb")
	}
	return borrado
}

func borrarRecursiva[K comparable, V any](nodo *nodoAb[K, V], clave K, comparar func(K, K) int, cantidad *int) (*nodoAb[K, V], V, bool) {
	if nodo == nil {
		var cero V
		return nil, cero, false
	}
	cmp := comparar(clave, nodo.clave)
	if cmp < 0 {
		var elem V
		var ok bool
		nodo.izq, elem, ok = borrarRecursiva(nodo.izq, clave, comparar, cantidad)
		return nodo, elem, ok
	} else if cmp > 0 {
		var elem V
		var ok bool
		nodo.der, elem, ok = borrarRecursiva(nodo.der, clave, comparar, cantidad)
		return nodo, elem, ok
	}

	borrado := nodo.dato
	if cantidad != nil {
		(*cantidad)--
	}

	if nodo.izq == nil {
		return nodo.der, borrado, true
	}
	if nodo.der == nil {
		return nodo.izq, borrado, true
	}

	sucesor := obtenerMinimo(nodo.der)
	nodo.clave = sucesor.clave
	nodo.dato = sucesor.dato

	var _ V
	nodo.der, _, _ = borrarRecursiva(nodo.der, sucesor.clave, comparar, nil)

	return nodo, borrado, true
}

func insertarYActualizar[K comparable, V any](n *nodoAb[K, V], clave K, dato V, cmp func(K, K) int) (*nodoAb[K, V], bool) {
	if n == nil {
		return &nodoAb[K, V]{clave: clave, dato: dato}, true
	}
	comp := cmp(clave, n.clave)
	if comp < 0 {
		var creado bool
		n.izq, creado = insertarYActualizar(n.izq, clave, dato, cmp)
		return n, creado
	} else if comp > 0 {
		var creado bool
		n.der, creado = insertarYActualizar(n.der, clave, dato, cmp)
		return n, creado
	} else {
		n.dato = dato
		return n, false
	}

}

func (a *ab[K, V]) Guardar(clave K, dato V) {
	creado := false
	a.raiz, creado = insertarYActualizar(a.raiz, clave, dato, a.comparar)
	if creado {
		a.cant++
	}
}

func obtenerRec[K comparable, V any](n *nodoAb[K, V], clave K, cmp func(K, K) int) *nodoAb[K, V] {
	if n == nil {
		return nil
	}
	comp := cmp(clave, n.clave)
	if comp < 0 {
		return obtenerRec(n.izq, clave, cmp)
	}
	if comp > 0 {
		return obtenerRec(n.der, clave, cmp)
	}
	return n
}
func (a *ab[K, V]) Obtener(clave K) V {
	n := obtenerRec(a.raiz, clave, a.comparar)
	if n == nil {
		panic("La clave no pertenece al abb")
	}
	return n.dato
}

func (a *ab[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	iterarInOrder(a.raiz, visitar)
}

func iterarInOrder[K comparable, V any](nodo *nodoAb[K, V], f func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}
	if !iterarInOrder(nodo.izq, f) {
		return false
	}
	if !f(nodo.clave, nodo.dato) {
		return false
	}
	return iterarInOrder(nodo.der, f)
}

//iterador externo

func (a *ab[K, V]) Iterador() IterDiccionario[K, V] {
	return nuevoIteradorExternoABB(a.raiz)
}

func nuevoIteradorExternoABB[K comparable, V any](raiz *nodoAb[K, V]) *iteradorExternoABB[K, V] {
	it := &iteradorExternoABB[K, V]{pila: TDAPila.CrearPilaDinamica[*nodoAb[K, V]]()}
	it.apilarIzquierda(raiz)
	return it
}

func (it *iteradorExternoABB[K, V]) apilarIzquierda(n *nodoAb[K, V]) {
	for n != nil {
		it.pila.Apilar(n)
		n = n.izq
	}
}

func (it *iteradorExternoABB[K, V]) HaySiguiente() bool {
	return !it.pila.EstaVacia()
}

func (it *iteradorExternoABB[K, V]) VerActual() (K, V) {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := it.pila.VerTope()
	return nodo.clave, nodo.dato
}

func (it *iteradorExternoABB[K, V]) Siguiente() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := it.pila.Desapilar()
	it.apilarIzquierda(nodo.der)

}

// iterador rango

func (a *ab[K, V]) IterarRango(desde, hasta *K, visitar func(K, V) bool) {
	iterarRangoRecursivamente(a.raiz, desde, hasta, a.comparar, visitar)
}

func iterarRangoRecursivamente[K comparable, V any](nodo *nodoAb[K, V], desde, hasta *K, comparar func(K, K) int, visitar func(K, V) bool) bool {
	if nodo == nil {
		return true
	}
	if desde != nil && comparar(nodo.clave, *desde) < 0 {
		return iterarRangoRecursivamente(nodo.der, desde, hasta, comparar, visitar)
	}
	if hasta != nil && comparar(nodo.clave, *hasta) > 0 {
		return iterarRangoRecursivamente(nodo.izq, desde, hasta, comparar, visitar)
	}
	if !iterarRangoRecursivamente(nodo.izq, desde, hasta, comparar, visitar) {
		return false
	}
	if !visitar(nodo.clave, nodo.dato) {
		return false
	}
	return iterarRangoRecursivamente(nodo.der, desde, hasta, comparar, visitar)
}

func (a *ab[K, V]) IteradorRango(desde, hasta *K) IterDiccionario[K, V] {
	it := &iteradorRangoABB[K, V]{pila: TDAPila.CrearPilaDinamica[*nodoAb[K, V]](), desde: desde, hasta: hasta, comparar: a.comparar}
	it.inicializarPilaRango(a.raiz)
	return it
}

func (it *iteradorRangoABB[K, V]) inicializarPilaRango(n *nodoAb[K, V]) {
	for n != nil {
		if it.desde != nil && it.comparar(n.clave, *it.desde) < 0 {
			n = n.der
		} else {
			it.pila.Apilar(n)
			n = n.izq
		}
	}
}

func (it *iteradorRangoABB[K, V]) HaySiguiente() bool {
	return !it.pila.EstaVacia()
}

func (it *iteradorRangoABB[K, V]) VerActual() (K, V) {
	if !it.HaySiguiente() {
		panic("El iterador rango terminó de iterar")
	}
	nodo := it.pila.VerTope()
	return nodo.clave, nodo.dato
}

func (it *iteradorRangoABB[K, V]) Siguiente() {
	if !it.HaySiguiente() {
		panic("El iterador rango terminó de iterar")
	}
	nodo := it.pila.Desapilar()

	nodoActual := nodo.der
	for nodoActual != nil {
		if it.hasta != nil && it.comparar(nodoActual.clave, *it.hasta) > 0 {
			nodoActual = nodoActual.izq
		} else {
			it.pila.Apilar(nodoActual)
			nodoActual = nodoActual.izq
		}
	}
}
