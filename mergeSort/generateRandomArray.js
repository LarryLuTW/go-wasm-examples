function _shuffle(array) {
  let tmp, current
  let top = array.length

  while (--top) {
    current = Math.floor(Math.random() * (top + 1))
    tmp = array[current]
    array[current] = array[top]
    array[top] = tmp
  }

  return array
}

function generateRandomArray() {
  // https://stackoverflow.com/questions/5836833/create-an-array-with-random-values
  let a = []
  for (i = 0; i < 1000000; ++i) a[i] = i
  a = _shuffle(a)
  return a
}
