package twosum

// Raciocínio para resolver de novo:
//  1. O par precisa somar target e cada índice só pode ser usado uma vez.
//  2. Dois laços acham o par, mas repetem trabalho: para cada número, o outro é sempre target - n.
//  3. Se os números já vistos estiverem numa tabela (valor → índice), o complemento se consulta na hora.
//  4. Consultar antes de guardar o número atual evita usar o mesmo índice duas vezes.
//  5. A resposta são os dois índices, o já visto primeiro.
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return nil
}
