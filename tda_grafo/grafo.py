from typing import Dict, Any
import random

PANIC_VERTICE_INEXISTENTE = "El Vertice no existe"
PANIC_VERTICES_INEXISTENTES = "Alguno de los Vertices no existe"
PANIC_ARISTA_INEXISTENTE = "La arista no existe"
PANIC_GRAFO_VACIO = "Aun no ha Vertices"


class Grafo:
    def __init__(self, dirigido: bool = False):
        self.esdirigido = dirigido
        self.lista_adyacencia = {}

    def agregar_vertice(self, v):
        if v not in self.lista_adyacencia:
            self.lista_adyacencia[v] = {}

    def borrar_vertice(self, a_borrar):
        if a_borrar not in self.lista_adyacencia:
            raise (PANIC_VERTICE_INEXISTENTE)

        if self.esdirigido:
            for vertice in list(self.lista_adyacencia.keys()):
                if a_borrar in self.lista_adyacencia[vertice]:
                    del self.lista_adyacencia[vertice][a_borrar]
        else:
            for vecino in list(self.lista_adyacencia[a_borrar].keys()):
                if vecino in self.lista_adyacencia:
                    self.lista_adyacencia[vecino].pop(a_borrar, None)
        del self.lista_adyacencia[a_borrar]

    def agregar_arista(self, v, w, peso: float = 1.0):
        if v not in self.lista_adyacencia or w not in self.lista_adyacencia:
            raise (PANIC_VERTICES_INEXISTENTES)
        self.lista_adyacencia[v][w] = peso

        if not self.esdirigido:
            self.lista_adyacencia[w][v] = peso

    def borrar_arista(self, v, w):
        if not self.estan_unidos(v, w):
            raise (PANIC_ARISTA_INEXISTENTE)
        del self.lista_adyacencia[v][w]
        if not self.esdirigido:
            del self.lista_adyacencia[w][v]

    def estan_unidos(self, v, w):
        return v in self.lista_adyacencia and w in self.lista_adyacencia[v]

    def peso_arista(self, v, w):
        if not self.estan_unidos(v, w):
            raise (PANIC_ARISTA_INEXISTENTE)
        return self.lista_adyacencia[v][w]

    def obtener_vertices(self):
        return list(self.lista_adyacencia.keys())

    def hay_vertice(self, v):
        return v in self.lista_adyacencia

    def adyacentes(self, v):
        if v not in self.lista_adyacencia:
            raise (PANIC_VERTICE_INEXISTENTE)

        return list(self.lista_adyacencia[v].keys())

    def vertice_aleatorio(self):
        if not self.lista_adyacencia:
            raise (PANIC_GRAFO_VACIO)
        return random.choice(list(self.lista_adyacencia.keys()))
