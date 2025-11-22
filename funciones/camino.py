from funciones_auxiliares import reconstruir_camino
import tda_grafo.grafo as g
from collections import deque


def camino_minimo_origen_destino(origen, destino, grafo: g.Grafo):
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

        for w in grafo.adyacentes():
            if w not in visitados:
                visitados.add(w)
                padres[w] = v
                cola.append(w)

    return None
