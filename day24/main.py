from z3 import *
import sys

z3.set_param("parallel.enable", True)

s = z3.SolverFor("QF_BV")
lines = sys.stdin.readlines()

values = {
	'x': z3.BitVecVal(0, 32),
	'y': z3.BitVecVal(0, 32),
	'z': z3.BitVecVal(0, 32),
	'w': z3.BitVecVal(0, 32),
}

inputs = []

def get(arg):
	if arg in values:
		return values[arg]
	else:
		return z3.BitVecVal(int(arg), 32)

for line in lines:
	line = line.strip()
	parts = line.split(' ')

	if parts[0] == 'inp':
		target = parts[1]

		idx = len(inputs)
		name = "inp" + str(idx)
		input = z3.BitVec(name, 32)
		inputs.append(input)
		s.add(z3.And(1 <= input, input <= 9))

		values[target] = input

	elif parts[0] == 'add':
		target = parts[1]
		arg = parts[2]

		values[target] = z3.simplify(values[target] + get(arg))

	elif parts[0] == 'mul':
		target = parts[1]
		arg = parts[2]

		values[target] = z3.simplify(values[target] * get(arg))

	elif parts[0] == 'div':
		target = parts[1]
		arg = parts[2]

		values[target] = z3.simplify(values[target] / get(arg))

	elif parts[0] == 'mod':
		target = parts[1]
		arg = parts[2]

		values[target] = z3.simplify(values[target] % get(arg))

	elif parts[0] == 'eql':
		target = parts[1]
		arg = parts[2]

		values[target] = z3.simplify(z3.If(values[target] == get(arg), z3.BitVecVal(1, 32), z3.BitVecVal(0, 32)))

# print(values['z'])
s.add(values['z'] == 0)
print(s.check())

ans = ""

for input in inputs:
	here = None
	# for j in range(1, 10):
	for j in range(9, 0, -1):
		s.push()
		s.add(input == z3.BitVecVal(j, 32))
		result = s.check()
		good = result == z3.sat
		s.pop()
		if good:
			s.add(input == z3.BitVecVal(j, 32))
			here = j
			break
	if here is None:
		print("unsat")
	ans += str(here)
	print(ans)

# print(inputs)

