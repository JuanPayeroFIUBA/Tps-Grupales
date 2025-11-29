from collections import deque
from tda_grafo.grafo import Grafo


# Funcion Auxiliar encargada de reconstuir el camino entre dos vertices a traves de un diccionario de padres
def reconstruir_camino(padres, origen, destino):
    actual = destino
    salida = []

    while actual != origen:
        salida.append(actual)
        actual = padres[actual]
    salida.append(origen)
    salida.reverse()
    return salida


# Funcion Auxiliar encargada de obtener los grados de entrada de un grafo y almacenando los adyacentes de cada vertice
def obtener_grados_de_entrada_subgrafo(vertices, grafo: Grafo):
    grados_ent = {}
    adyacentes = {}
    for v in vertices:
        grados_ent[v] = 0
        adyacentes[v] = []

    for v in vertices:
        for w in grafo.adyacentes(v):
            if w in vertices:
                grados_ent[w] += 1
                adyacentes[v].append(w)
    return grados_ent, adyacentes


# Funcion que crea un subgrafo que tiene parte de los vertices del pasado por parametro, y las aristas de estos invertidas
def crear_subgrafo(vertices, grafo: Grafo):
    nuevo = Grafo(True)

    for v in vertices:
        if not grafo.hay_vertice(v):
            return None
        nuevo.agregar_vertice(v)

    for v in vertices:
        for w in grafo.adyacentes(v):
            if w in vertices:
                nuevo.agregar_arista(w, v)
    return nuevo


# Funcion que encuentra el camino minimo (por BFS) desde un vertice hasta todos los demas del grafo, retornando el camino al vertice mas lejano (reconstruido a partir de un diccionario de padres) y su distancia
def camino_minimo_grafo(origen, grafo: Grafo):
    cola = deque()
    padres = {}
    visitados = set()
    orden = {}
    maximo_orden_encolado = 0  # la distancia minima maxima hasta ahora
    maximo_vertice = ""

    cola.append(origen)

    padres[origen] = None
    orden[origen] = 0
    visitados.add(origen)

    while cola:
        v = cola.popleft()

        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                padres[w] = v
                cola.append(w)
                orden[w] = orden[v] + 1

                if orden[w] > maximo_orden_encolado:
                    maximo_orden_encolado = orden[w]
                    maximo_vertice = w

    return reconstruir_camino(padres, origen, maximo_vertice), maximo_orden_encolado


# Funcion que realiza un recorrido DFS sobre un grafo, en busca de ciclos de una longitud pasada por parametro, y guardando el recorrido en un diccionario de vertices, si no hubiera tal recorrido de ese largo, se retornaria None
def dfs_ciclo(origen, actual, largo, recorrido, visitados, grafo: Grafo):
    if len(recorrido) == largo:
        if origen in grafo.adyacentes(actual):
            recorrido.append(origen)
            return recorrido
        return None

    for w in grafo.adyacentes(actual):
        if w == origen or w in visitados:
            continue

        visitados.add(w)
        recorrido.append(w)

        camino = dfs_ciclo(origen, w, largo, recorrido, visitados, grafo)
        if camino is not None:
            return camino

        visitados.remove(w)
        recorrido.pop()

    return None
