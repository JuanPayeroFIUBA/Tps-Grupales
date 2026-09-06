#!/usr/bin/python3
import sys
from funciones.ciclo import obtener_ciclo
from funciones.diametro import obtener_diametro
from funciones.camino import camino_minimo_origen_destino
from funciones.lectura import orden_topologico_lectura
from funciones.rango import obtener_paginas_en_rango
from funciones.navegacion import navegar_primer_link
import tda_grafo.grafo as g

COMANDO_LISTAR = "listar_operaciones"
COMANDO_CAMINO = "camino"
COMANDO_CICLOS = "ciclo"
COMANDO_LECTURA = "lectura"
COMANDO_DIAMETRO = "diametro"
COMANDO_RANGO = "rango"
COMANDO_NAVEGACION = "navegacion"

LISTA_COMANDOS = f"""
{COMANDO_CAMINO}
{COMANDO_CICLOS}
{COMANDO_LECTURA}
{COMANDO_DIAMETRO}
{COMANDO_RANGO}
{COMANDO_NAVEGACION}
""".strip()

ERROR_CAMINO_NO_ENCONTRADO = "No se encontro recorrido"
ERROR_LECTURA = "No existe forma de leer las paginas en orden"
ERROR_COMANDO_INVALIDO = "Comando invalido"


def procesar_salida(error_por_pantalla, datos, hay_que_imprimir_costo=False):
    if not datos:
        print(error_por_pantalla)
    else:
        print(" -> ".join(datos))
        if hay_que_imprimir_costo:
            print(f"Costo: {len(datos) - 1}")


def main():
    if len(sys.argv) < 2:
        print("Uso: ./netstats <archivo_wiki>")
        return

    ruta_archivo = sys.argv[1]
    grafo = crear_grafo_desde_archivo(ruta_archivo)
    procesar_entradas(grafo)
    return


def crear_grafo_desde_archivo(ruta):
    grafo = g.Grafo(dirigido=True)

    with open(ruta, "r", encoding="utf-8") as archivo:
        for linea in archivo:
            partes = linea.rstrip("\n").split("\t")
            if not partes:
                continue

            pagina = partes[0]
            grafo.agregar_vertice(pagina)

            for link in partes[1:]:
                link = link.strip()
                if not link:
                    continue

                grafo.agregar_vertice(link)
                grafo.agregar_arista(pagina, link)

    return grafo


def procesar_entradas(grafo: g.Grafo):
    for linea in sys.stdin:
        linea = linea.strip()
        if not linea:
            continue

        partes = linea.split(" ", 1)
        comando = partes[0]
        parametros = partes[1] if len(partes) > 1 else ""

        if comando == COMANDO_LISTAR:
            print(LISTA_COMANDOS)

        elif comando == COMANDO_CAMINO:
            origen, destino = parametros.split(",")

            camino = camino_minimo_origen_destino(origen, destino, grafo)
            procesar_salida(ERROR_CAMINO_NO_ENCONTRADO, camino, True)

        elif comando == COMANDO_LECTURA:
            paginas = parametros.split(",")
            vertices = set(paginas)

            camino = orden_topologico_lectura(vertices, grafo)
            procesar_salida(ERROR_LECTURA, camino)

        elif comando == COMANDO_DIAMETRO:
            camino, diametro = obtener_diametro(grafo)
            print(" -> ".join(camino))
            print(f"Costo: {diametro}")

        elif comando == COMANDO_CICLOS:
            partes_ciclo = parametros.split(",")
            origen = partes_ciclo[0]
            largo = int(partes_ciclo[1])

            if not grafo.existe_vertice(origen):
                print(ERROR_CAMINO_NO_ENCONTRADO)
                continue

            ciclo = obtener_ciclo(origen, largo, grafo)

            procesar_salida(ERROR_CAMINO_NO_ENCONTRADO, ciclo)

        elif comando == COMANDO_RANGO:
            partes_rango = parametros.split(",")
            pagina = partes_rango[0]
            n = int(partes_rango[1])

            if not grafo.existe_vertice(pagina):
                print("0")
                continue

            cantidad = obtener_paginas_en_rango(pagina, n, grafo)
            print(cantidad)

        elif comando == COMANDO_NAVEGACION:
            pagina_origen = parametros.strip()

            if not grafo.existe_vertice(pagina_origen):
                print(pagina_origen)
                continue

            recorrido = navegar_primer_link(pagina_origen, grafo)
            print(" -> ".join(recorrido))

        else:
            print(ERROR_COMANDO_INVALIDO)


if __name__ == "__main__":
    main()
