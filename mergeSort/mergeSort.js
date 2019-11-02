function merge(left, right) {
  let i = 0
  let j = 0
  let k = 0

  let arr = []
  while (k < left.length + right.length) {
    if (i > left.length - 1 && j <= right.length - 1) {
      arr.push(right[j])
      j++
    } else if (j > right.length - 1 && i <= left.length - 1) {
      arr.push(left[i])
      i++
    } else if (left[i] < right[j]) {
      arr.push(left[i])
      i++
    } else {
      arr.push(right[j])
      j++
    }
    k++
  }
  return arr
}

function _mergeSortJS(arr) {
  if (arr.length < 2) {
    return arr
  }
  let mid = arr.length / 2

  let left = _mergeSortJS(arr.slice(0, mid))
  let right = _mergeSortJS(arr.slice(mid, arr.length))

  return merge(left, right)
}

function mergeSortJS(arr) {
  console.time('MergeSort in JS')
  let result = _mergeSortJS(arr)
  console.timeEnd('MergeSort in JS')
  return result
}
