from funciones.funciones_auxiliares import dfs_ciclo
from tda_grafo.grafo import Grafo


# Funcion que ejecuta el comando "ciclo" que permite obtener algun ciclo de paginas de largo n que comience en la página indicada
def obtener_ciclo(origen, largo, grafo: Grafo):
    if largo > len(grafo):
        return None

    visitados = set()
    camino = []

    camino.append(origen)
    visitados.add(origen)

    return dfs_ciclo(origen, origen, largo, camino, visitados, grafo)
