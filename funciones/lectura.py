from collections import deque
from funciones.funciones_auxiliares import (
    obtener_grados_de_entrada_subgrafo,
    crear_subgrafo,
)
from tda_grafo.grafo import Grafo


def orden_topologico_lectura(vertices, grafo):
    nuevo = crear_subgrafo(vertices, grafo)
    if nuevo is None:
        return None
    g_ent, adyacentes = obtener_grados_de_entrada_subgrafo(vertices, nuevo)

    cola = deque()
    camino_topologico = []

    for v in vertices:
        if g_ent[v] == 0:
            cola.append(v)

    while cola:
        v = cola.popleft()
        camino_topologico.append(v)
        for w in adyacentes[v]:
            g_ent[w] -= 1
            if g_ent[w] == 0:
                cola.append(w)

    if len(camino_topologico) != len(vertices):
        return None
    return camino_topologico
