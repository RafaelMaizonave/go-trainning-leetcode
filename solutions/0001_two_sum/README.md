# 1. Two Sum

- Dificuldade: Easy
- Tags: array, lookup table
- Link: https://leetcode.com/problems/two-sum/
- Status: Resolvido

## Notas

Um par só precisa ser encontrado uma vez. Guardar o valor já visto e o índice evita percorrer o array de novo para cada elemento.

## Resultado

- Abordagem: tabela de consulta. Para cada número, verifico se o complemento (target - número) já apareceu.
- Complexidade: O(n) tempo e O(n) memória.
