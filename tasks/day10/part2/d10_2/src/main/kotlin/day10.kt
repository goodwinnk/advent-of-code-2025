package nk2025.day10

import org.ssclab.log.SscLogger
import org.ssclab.pl.milp.*
import java.io.File
import kotlin.math.roundToInt

private fun IntArray.toDoubleArray(): DoubleArray =
    this.map { it.toDouble() }.toDoubleArray()

private data class MachineRaw(
    val on: String,
    val buttons: List<List<Int>>,
    val joltage: List<Int>,
)

private fun parse(input: String): List<MachineRaw> {
    val result = mutableListOf<MachineRaw>()

    for (rawLine in input.lines()) {
        val line = rawLine.trim()
        if (line.isEmpty()) continue

        // on section: [ ... ]
        val iL = line.indexOf('[')
        val iR = line.indexOf(']')
        if (iL == -1 || iR == -1 || iL > iR) {
            throw IllegalArgumentException("invalid line (on): $line")
        }
        val on = line.substring(iL + 1, iR).trim()

        // joltage section: { ... } (last pair)
        val jL = line.lastIndexOf('{')
        val jR = line.lastIndexOf('}')
        if (jL == -1 || jR == -1 || jL > jR || jL < iR) {
            throw IllegalArgumentException("invalid line (joltage): $line")
        }
        val joltage = parseInts(line.substring(jL + 1, jR))

        // buttons section(s) between ']' and '{': ( ... )( ... )...
        var middle = line.substring(iR + 1, jL)
        val buttons = mutableListOf<List<Int>>()
        while (true) {
            val bL = middle.indexOf('(')
            if (bL == -1) break
            val bRRel = middle.substring(bL + 1).indexOf(')')
            if (bRRel == -1) throw IllegalArgumentException("unclosed buttons group: $line")
            val bR = bL + 1 + bRRel

            val group = middle.substring(bL + 1, bR)
            buttons += parseInts(group)

            middle = middle.substring(bR + 1)
        }

        result += MachineRaw(on = on, buttons = buttons, joltage = joltage)
    }

    return result
}

private fun parseInts(s: String): List<Int> =
    s.split(',').map(String::toInt)

private fun MachineRaw.solve(): Solution? {
    val c = IntArray(this.buttons.size) { 1 }

    val a: Array<IntArray> = Array(this.joltage.size) { j ->
        IntArray(this.buttons.size)
    }
    for ((i, button) in this.buttons.withIndex()) {
        for (br in button) {
            a[br][i] = 1
        }
    }

    return solve(c, a, this.joltage.toIntArray())
}

private fun solve(c: IntArray, a: Array<IntArray>, b: IntArray): Solution? {
    val f = LinearObjectiveFunction(c.toDoubleArray(), GoalType.MIN)

    val constraints = ArrayList<Constraint>()
    for (i in a.indices) {
        constraints.add(Constraint(a[i].toDoubleArray(), ConsType.EQ, b[i].toDouble()))
    }
    constraints.add(Constraint(DoubleArray(c.size) { 1.0 }, ConsType.INT, Double.NaN, "Integers"))

    val lp = MILP(f, constraints)
    val solutionType = lp.resolve()

    if (solutionType == SolutionType.OPTIMAL) {
        val solution = lp.getSolution()
        return solution
    } else {
        SscLogger.log("no optimal solution:$solutionType")
        return null
    }
}

private fun part2(input: String): Int {
    val machines = parse(input)
    return machines.sumOf { machine ->
        println(machine)
        val solution = machine.solve() ?: error("Couldn't solve machine: $machine")
        val solutionInt = solution.optimumValue.roundToInt()
        println("$solutionInt ${solution.optimumValue} ${solution.valuesSolution.toList()}")
        solutionInt
    }
}

fun main() {
    println("Day 10")
    println(part2(
        """
            [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
            [...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
            [.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
        """.trimIndent()
    ))
    println(part2(File("../../../../data/day10/task.txt").readText()))
}