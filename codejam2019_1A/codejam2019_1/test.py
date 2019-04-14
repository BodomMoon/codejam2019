def odd(r, C, s):
for c in xrange(s, C+1):
yield (r, c)
for c in xrange(1, s):
yield (r, c)

def even(r, C, s):
for c in xrange(1, s):
yield (r, c)
for c in reversed(xrange(s, C+1)):
yield (r, c)

def pylons():
R, C = map(int, raw_input().strip().split())
swapped = False
if R > C:
R, C = C, R
swapped = True
if R == 2 and C < 5:
return "IMPOSSIBLE"
if R == 3 and C == 3:
return "IMPOSSIBLE"

result = []
r = 0
if R == 4 and C == 4:
iter1 = odd(r+1, C, 1)
iter2 = odd(r+2, C, 3)
iter3 = odd(r+3, C, 1)
iter4 = odd(r+4, C, 3)
for c in xrange(C):
result.append(next(iter1))
result.append(next(iter2))
result.append(next(iter3))
result.append(next(iter4))
result[12], result[14] = result[14], result[12]
r += 4
else:
if R % 2 != 0 and R >= 3:
iter1 = odd(r+1, C, 1)
iter2 = odd(r+2, C, 3)
iter3 = odd(r+3, C, 1)
for c in xrange(C):
result.append(next(iter1))
result.append(next(iter2))
result.append(next(iter3))
r += 3

if (R-r) % 2 == 0 and C >= 5:
while r != R:
if not result or abs(r+2 - result[-1][0]) != abs(3 - result[-1][1]):
iter1 = odd(r+2, C, 3)
iter2 = even(r+1, C, C-1)
else:
iter1 = odd(r+1, C, 3)
iter2 = even(r+2, C, C-1)
for c in xrange(C):
result.append(next(iter1))
result.append(next(iter2))
r += 2

if swapped:
swapped = False
for i in xrange(len(result)):
result[i] = (result[i][1], result[i][0])
R, C = C, R

assert(R*C == len(result))
for i in xrange(1, len(result)):
assert(abs(result[i][0]-result[i-1][0]) != abs(result[i][1]-result[i-1][1]) and
result[i][0]-result[i-1][0] != 0 and result[i][1]-result[i-1][1] != 0)
return "POSSIBLE\n{}".format("\n".join(map(lambda x: " ".join(map(str, x)), result)))

for case in xrange(input()):
print 'Case #%d: %s' % (case+1, pylons())