from funciones.funciones_auxiliares import camino_minimo_grafo
from tda_grafo.grafo import Grafo


# funcion que ejecuta el comando "diametro", que permite obtener el diámetro de toda la red, osea obtener el camino mínimo más grande de toda la red.
def obtener_diametro(grafo: Grafo):
    camino_minimo_maximo = []
    maximo = 0
    for v in grafo.obtener_vertices():
        camino, largo = camino_minimo_grafo(v, grafo)
        if largo > maximo:
            camino_minimo_maximo = camino
            maximo = largo

    return camino_minimo_maximo, maximo
