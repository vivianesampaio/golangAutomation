# language: pt
Funcionalidade: Calculadora

  Esquema do Cenario: Calcular valores
    Quando realizo a operação <v1> <operação> <v2>
    Então Então o valor deve ser <resultado>

    Exemplos:
      | v1 | v2 | operação | resultado |
      |  2 |  2 | +        |         4 |
      |  2 |  2 | -        |         0 |
      |  2 |  2 | /        |         1 |
      |  2 |  2 | *        |         4 |
