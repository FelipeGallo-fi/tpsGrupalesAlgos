// Implementación del Árbol Binario de Búsqueda (ABB) genérico.
// Se utiliza como base para ordenar vuelos por fecha.
// El archivo abb_fechas_vuelos.go se apoya en esta implementación.

package diccionario

import (
	TDAPila "tp2/tdas/abb/pila"
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

func buscarClave[K comparable, V any](raiz *nodoAb[K, V], clave K, comparar func(K, K) int) (padre *nodoAb[K, V], nodo *nodoAb[K, V]) {
	actual := raiz
	var padreAux *nodoAb[K, V]

	for actual != nil {
		cmp := comparar(clave, actual.clave)
		if cmp == 0 {
			return padreAux, actual
		} else if cmp < 0 {
			padreAux = actual
			actual = actual.izq
		} else {
			padreAux = actual
			actual = actual.der
		}
	}
	return padreAux, nil
}

func (a *aBB[K, V]) Pertenece(clave K) bool {
	_, nodo := buscarClave(a.raiz, clave, a.comparar)
	return nodo != nil
}

func (a *aBB[K, V]) Cantidad() int {
	return a.cant
}

func (a *aBB[K, V]) Borrar(clave K) V {
	padre, nodo := buscarClave(a.raiz, clave, a.comparar)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	borrado := nodo.dato

	if nodo.izq == nil || nodo.der == nil {
		a.borrarNodoConUnHijo(padre, nodo)
	} else {
		a.borrarNodoConDosHijos(nodo)
	}

	a.cant--
	return borrado
}

func (a *aBB[K, V]) borrarNodoConUnHijo(padre, nodo *nodoAb[K, V]) {
	var hijo *nodoAb[K, V]
	if nodo.izq != nil {
		hijo = nodo.izq
	} else {
		hijo = nodo.der
	}
	a.reemplazarEnPadre(padre, nodo, hijo)
}

func (a *aBB[K, V]) borrarNodoConDosHijos(nodo *nodoAb[K, V]) {
	sucesorPadre := nodo
	sucesor := nodo.der
	for sucesor.izq != nil {
		sucesorPadre = sucesor
		sucesor = sucesor.izq
	}
	nodo.clave = sucesor.clave
	nodo.dato = sucesor.dato
	a.reemplazarEnPadre(sucesorPadre, sucesor, sucesor.der)
}

func (a *aBB[K, V]) reemplazarEnPadre(padre, nodo, nuevo *nodoAb[K, V]) {
	if padre == nil {
		a.raiz = nuevo
	} else if padre.izq == nodo {
		padre.izq = nuevo
	} else {
		padre.der = nuevo
	}
}

func (a *aBB[K, V]) Guardar(clave K, dato V) {
	if a.raiz == nil {
		a.raiz = &nodoAb[K, V]{clave: clave, dato: dato}
		a.cant++
		return
	}

	padre, nodo := buscarClave(a.raiz, clave, a.comparar)
	if nodo != nil {
		nodo.dato = dato
		return
	}

	nuevo := &nodoAb[K, V]{clave: clave, dato: dato}
	a.reemplazarEnPadre(padre, nil, nuevo)
	a.cant++
}

func (a *aBB[K, V]) Obtener(clave K) V {
	_, nodo := buscarClave(a.raiz, clave, a.comparar)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

func (a *aBB[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	iterarRangoRec(a.raiz, nil, nil, a.comparar, visitar)
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
