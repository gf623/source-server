var fa = [];
for (var i = 0; i < 10; i++) {
    fa[i] = i
}

var find = (i) => {
    if(fa[i] == i) {
        return i
    } else {
        return find(fa[i])
    }
}

var merge = (i, j) => {
    fa[find(i)] = find(j)
}

merge(1, 2)
merge(1, 3)
merge(1, 4)
merge(1, 5)
merge(1, 6)
merge(1, 7)
merge(1, 8)
merge(1, 9)
// merge(4, 5)
console.log('__arr:', fa, find(3), find(5))