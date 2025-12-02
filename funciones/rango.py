from collections import deque
from tda_grafo.grafo import Grafo

def obtener_paginas_en_rango(pagina, n, grafo: Grafo):
    cola = deque()
    visitados = set()
    distancias = {}

    cola.append(pagina)
    visitados.add(pagina)
    distancias[pagina] = 0

    while cola:
        v = cola.popleft()

        if distancias[v] > n:
            break  # si ya pasamos el nivel frenamos

        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                distancias[w] = distancias[v] + 1
                cola.append(w)

    contador = 0
    for vertice, distancia in distancias.items():
        if distancia == n:
            contador += 1

    return contador
