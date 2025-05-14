package diccionario

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

func (a *ab[K, V]) panicPertenece(clave K) {
	if !a.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
}

func (a *ab[K, V]) Borrar(clave K) V {
	a.panicPertenece(clave)
	var elementoBorrado V
	a.raiz, elementoBorrado = borrarRecursiva(a.raiz, clave, a.comparar, &a.cant)
	return elementoBorrado
}

func obtenerMinimo[K comparable, V any](nodo *nodoAb[K, V]) *nodoAb[K, V] {
	if nodo == nil || nodo.izq == nil {
		return nodo
	}
	return obtenerMinimo(nodo.izq)
}

func borrarRecursiva[K comparable, V any](nodo *nodoAb[K, V], clave K, comparar func(K, K) int, cantidad *int) (*nodoAb[K, V], V) {
	if nodo == nil {
		var cero V
		return nil, cero
	}
	cmp := comparar(clave, nodo.clave)

	if cmp < 0 {
		var dato V
		nodo.izq, dato = borrarRecursiva(nodo.izq, clave, comparar, cantidad)
		return nodo, dato
	} else if cmp > 0 {
		var dato V
		nodo.der, dato = borrarRecursiva(nodo.der, clave, comparar, cantidad)
		return nodo, dato
	} else {
		*cantidad--

		if nodo.izq == nil {
			return nodo.der, nodo.dato
		}

		if nodo.der == nil {
			return nodo.izq, nodo.dato
		}

		nuevoCantidato := obtenerMinimo(nodo.izq.der)
		nodo.clave = nuevoCantidato.clave
		nodo.dato = nuevoCantidato.dato
		nodo.der, _ = borrarRecursiva(nodo.der, nuevoCantidato.clave, comparar, cantidad)
		return nodo, nuevoCantidato.dato
	}

}

func insertarYActualizar[K comparable, V any](n *nodoAb[K, V], clave K, dato V, cmp func(K, K) int) (*nodoAb[K, V], bool) {
	if n == nil {
		return &nodoAb[K, V]{clave: clave, dato: dato}, true
	}
	var creado bool
	comp := cmp(clave, n.clave)
	if comp < 0 {
		n.izq, creado = insertarYActualizar(n.izq, clave, dato, cmp)
		return n, creado
	}
	if comp > 0 {
		n.der, creado = insertarYActualizar(n.der, clave, dato, cmp)
		return n, creado
	}
	n.dato = dato
	return n, false
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
		panic("La clave no pertenece al diccionario")
	}
	return n.dato
}

//---Iterador Interno----

func (a *ab[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	iterarRec(a.raiz, visitar)
}

func iterarRec[K comparable, V any](nodo *nodoAb[K, V], f func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}

	if !iterarRec(nodo.izq, f) {
		return false
	}

	if !f(nodo.clave, nodo.dato) {
		return false
	}

	return iterarRec(nodo.der, f)
}

//---Iterador Externo----

func (a *ab[K, V]) Iterador() *iteradorExternoABB[K, V] {
	return nuevoIteradorExternoABB(a.raiz)
}

type iteradorExternoABB[K comparable, V any] struct {
	pila []*nodoAb[K, V]
}

func nuevoIteradorExternoABB[K comparable, V any](raiz *nodoAb[K, V]) *iteradorExternoABB[K, V] {
	it := &iteradorExternoABB[K, V]{}
	it.apilarIzquierda(raiz)
	return it
}

func (it *iteradorExternoABB[K, V]) apilarIzquierda(n *nodoAb[K, V]) {
	for n != nil {
		it.pila = append(it.pila, n)
		n = n.izq
	}
}

func (it *iteradorExternoABB[K, V]) HaySiguiente() bool {
	return len(it.pila) > 0
}

func (it *iteradorExternoABB[K, V]) Siguiente() (clave K, dato V) {
	nodo := it.pila[len(it.pila)-1]
	return nodo.clave, nodo.dato
}

func (it *iteradorExternoABB[K, V]) Avanzar() {
	if len(it.pila) == 0 {
		return
	}

	nodo := it.pila[len(it.pila)-1]
	it.pila = it.pila[:len(it.pila)-1]
	it.apilarIzquierda(nodo.der)
}
