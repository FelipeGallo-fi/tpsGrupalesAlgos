package diccionario

type ab[K comparable, V any] struct{
	raiz *nodoAb[K, V]
	cant	int
	comparar func(K,K) 	int
}

type nodoAb[K comparable, V any] struct{
	izq *nodoAb[K, V]
	der *nodoAb[K, V]
	clave	K
	dato	V
}


func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V]{
	return &ab[K, V]{raiz: nil,cant: 0, comparar: funcion_cmp}
}



func (a *ab[K , V]) Pertenece(clave K) bool {
	return perteneceRecursiva(a.raiz,clave, a.comparar)
}

func perteneceRecursiva[K comparable, V any](nodo *nodoAb[K, V], clave K, comparar func(K, K) int) bool {
	if nodo == nil{
		return false
	}
	cmp := comparar(clave ,nodo.clave)
	if cmp == 0 {
		return true
	}else if cmp < 0 {
		return perteneceRecursiva(nodo.izq,clave , comparar)
	} else {
		return perteneceRecursiva(nodo.der ,clave ,comparar)
	}
}



func (a *ab[K, V]) Cantidad() int {
	return a.cant
}

func (a *ab[K, V]) panicPertenece(clave K) {
	if !a.Pertenece(clave){
		panic("La clave no pertenece al diccionario")
	}
}

func (a *ab[K, V]) Borrar(clave K) V {
	a.panicPertenece(clave)
	var elementoBorrado V
	a.raiz , elementoBorrado = borrarRecursiva(a.raiz,clave, a.comparar,&a.cant)
	return elementoBorrado
}

func obtenerMinimo[K comparable, V any](nodo *nodoAb[K ,V]) *nodoAb[K,V]{
	if nodo == nil || nodo.izq == nil {
		return nodo
	}
	return obtenerMinimo(nodo.izq)
}

func borrarRecursiva[K comparable, V any](nodo *nodoAb[K, V], clave K, comparar func(K, K) int, cantidad *int )(*nodoAb[K, V], V){
	if nodo == nil {
		var cero V
		return nil, cero
	}
	cmp := comparar(clave , nodo.clave)

	if cmp < 0 {
		var dato V
		nodo.izq, dato = borrarRecursiva(nodo.izq, clave , comparar, cantidad) 
		return nodo , dato
	} else if cmp > 0 {
		var dato V
		nodo.der, dato = borrarRecursiva(nodo.der,clave,comparar,cantidad)
		return nodo, dato
	} else {
		*cantidad --

		if nodo.izq == nil {
			return nodo.der , nodo.dato
		}

		if nodo.der == nil {
			return nodo.izq , nodo.dato
		}

		nuevoCantidato := obtenerMinimo(nodo.izq.der)
		nodo.clave = nuevoCantidato.clave
		nodo.dato =nuevoCantidato.dato
		nodo.der,_= borrarRecursiva(nodo.der,nuevoCantidato.clave, comparar , cantidad)
		return nodo ,nuevoCantidato.dato
	}
	
}