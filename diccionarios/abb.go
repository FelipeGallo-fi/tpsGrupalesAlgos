package diccionario

import (
	TDAPila "tdas/pila"
)

type aBB[K comparable, V any] struct {
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

type iteradorABB[K comparable, V any] struct {
	pila     TDAPila.Pila[*nodoAb[K, V]]
	desde    *K
	hasta    *K
	comparar func(K, K) int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &aBB[K, V]{raiz: nil, cant: 0, comparar: funcion_cmp}
}

func buscarNodoRec[K comparable, V any](n *nodoAb[K, V], clave K, comparar func(K, K) int) *nodoAb[K, V] {
	if n == nil {
		return nil
	}
	cmp := comparar(clave, n.clave)
	if cmp == 0 {
		return n
	} else if cmp < 0 {
		return buscarNodoRec(n.izq, clave, comparar)
	} else {
		return buscarNodoRec(n.der, clave, comparar)
	}
}

func (a *aBB[K, V]) Pertenece(clave K) bool {
	return buscarNodoRec(a.raiz, clave, a.comparar) != nil
}

func (a *aBB[K, V]) Cantidad() int {
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

func (a *aBB[K, V]) Borrar(clave K) V {
	var borrado V
	var borradoOk bool
	a.raiz, borrado, borradoOk = borrarRecursiva(a.raiz, clave, a.comparar, &a.cant)
	if !borradoOk {
		panic("La clave no pertenece al diccionario")
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

func (a *aBB[K, V]) Guardar(clave K, dato V) {
	creado := false
	a.raiz, creado = insertarYActualizar(a.raiz, clave, dato, a.comparar)
	if creado {
		a.cant++
	}
}

func (a *aBB[K, V]) Obtener(clave K) V {
	n := buscarNodoRec(a.raiz, clave, a.comparar)
	if n == nil {
		panic("La clave no pertenece al diccionario")
	}
	return n.dato
}

func (a *aBB[K, V]) Iterar(visitar func(clave K, dato V) bool) {
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

// iterador externo
func (a *aBB[K, V]) Iterador() IterDiccionario[K, V] {
	return a.IteradorRango(nil, nil)
}

func (a *aBB[K, V]) IteradorRango(desde, hasta *K) IterDiccionario[K, V] {
	return nuevoIteradorRangoABB(a.raiz, desde, hasta, a.comparar)
}

func nuevoIteradorRangoABB[K comparable, V any](raiz *nodoAb[K, V], desde, hasta *K, comparar func(K, K) int) *iteradorABB[K, V] {
	it := &iteradorABB[K, V]{pila: TDAPila.CrearPilaDinamica[*nodoAb[K, V]](), desde: desde, hasta: hasta, comparar: comparar}
	it.apilarNodos(raiz)
	return it
}

func (it *iteradorABB[K, V]) apilarNodos(n *nodoAb[K, V]) {
	for n != nil {
		if it.desde != nil && it.comparar(n.clave, *it.desde) < 0 {
			n = n.der
		} else if it.hasta != nil && it.comparar(n.clave, *it.hasta) > 0 {
			n = n.izq
		} else {
			it.pila.Apilar(n)
			n = n.izq
		}
	}
}

func (it *iteradorABB[K, V]) HaySiguiente() bool {
	return !it.pila.EstaVacia()
}

func (it *iteradorABB[K, V]) VerActual() (K, V) {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := it.pila.VerTope()
	return nodo.clave, nodo.dato
}

func (it *iteradorABB[K, V]) Siguiente() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := it.pila.Desapilar()
	it.apilarNodos(nodo.der)
}

func iterarRangoRec[K comparable, V any](n *nodoAb[K, V], desde, hasta *K, comparar func(K, K) int, visitar func(clave K, dato V) bool) bool {
	if n == nil {
		return true
	}

	if desde != nil && comparar(n.clave, *desde) < 0 {
		return iterarRangoRec(n.der, desde, hasta, comparar, visitar)
	}
	if hasta != nil && comparar(n.clave, *hasta) > 0 {
		return iterarRangoRec(n.izq, desde, hasta, comparar, visitar)
	}

	if !iterarRangoRec(n.izq, desde, hasta, comparar, visitar) {
		return false
	}
	if !visitar(n.clave, n.dato) {
		return false
	}
	return iterarRangoRec(n.der, desde, hasta, comparar, visitar)
}

func (a *aBB[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	iterarRangoRec(a.raiz, desde, hasta, a.comparar, visitar)
}
