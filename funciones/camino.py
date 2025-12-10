from funciones.funciones_auxiliares import reconstruir_camino
from tda_grafo.grafo import Grafo
from collections import deque


# Funcion encargada de ejecutar el comando "camino"
# que imprime una lista con las páginas con los cuales
# navegamos de la página origen a la página destino, navegando lo menos posible,
# para lo cual realiza un recorrido BFS para encontrar el camino minimo
def camino_minimo_origen_destino(origen, destino, grafo: Grafo):
    if not grafo.existe_vertice(origen) or not grafo.existe_vertice(destino):
        return None
    cola = deque()
    padres = {}
    visitados = set()

    cola.append(origen)
    padres[origen] = None
    visitados.add(origen)

    while cola:
        v = cola.popleft()

        if v == destino:
            return reconstruir_camino(padres, origen, destino)

        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                padres[w] = v
                cola.append(w)

    return None
