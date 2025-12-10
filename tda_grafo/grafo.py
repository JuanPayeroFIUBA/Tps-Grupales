import random
from typing import Any

MSJ_ERROR_VERTICE_INEXISTENTE = "El Vertice no existe"
MSJ_ERROR_VERTICES_INEXISTENTES = "Alguno de los Vertices no existe"
MSJ_ERROR_ARISTA_INEXISTENTE = "La arista no existe"
MSJ_ERROR_GRAFO_VACIO = "Aun no hay Vertices"


class Grafo:
    def __init__(self, dirigido: bool = False):
        self.esdirigido = dirigido
        self.lista_adyacencia = {}

    def agregar_vertice(self, v):
        if v not in self.lista_adyacencia:
            self.lista_adyacencia[v] = {}

    def borrar_vertice(self, v):
        if v not in self.lista_adyacencia:
            raise Exception(MSJ_ERROR_VERTICE_INEXISTENTE)

        for vertice in self.lista_adyacencia:
            self.lista_adyacencia[vertice].pop(v, None)
        del self.lista_adyacencia[v]

    def agregar_arista(self, v, w, peso: Any = 1):
        if v not in self.lista_adyacencia or w not in self.lista_adyacencia:
            raise Exception(MSJ_ERROR_VERTICES_INEXISTENTES)
        self.lista_adyacencia[v][w] = peso

        if not self.esdirigido:
            self.lista_adyacencia[w][v] = peso

    def borrar_arista(self, v, w):
        if not self.estan_unidos(v, w):
            raise Exception(MSJ_ERROR_ARISTA_INEXISTENTE)
        del self.lista_adyacencia[v][w]
        if not self.esdirigido:
            del self.lista_adyacencia[w][v]

    def estan_unidos(self, v, w):
        return v in self.lista_adyacencia and w in self.lista_adyacencia[v]

    def peso_arista(self, v, w):
        if not self.estan_unidos(v, w):
            raise Exception(MSJ_ERROR_ARISTA_INEXISTENTE)
        return self.lista_adyacencia[v][w]

    def obtener_vertices(self):
        return list(self.lista_adyacencia.keys())

    def existe_vertice(self, v):
        return v in self.lista_adyacencia

    def adyacentes(self, v):
        if v not in self.lista_adyacencia:
            raise Exception(MSJ_ERROR_VERTICE_INEXISTENTE)

        return list(self.lista_adyacencia[v].keys())

    def vertice_aleatorio(self):
        if not self.lista_adyacencia:
            raise Exception(MSJ_ERROR_GRAFO_VACIO)
        return random.choice(list(self.lista_adyacencia.keys()))

    def __len__(self):
        return len(self.lista_adyacencia)
