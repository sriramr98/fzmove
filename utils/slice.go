package utils

func SliceContains(slice []string, search string) bool {
  for _, val := range slice {
    if val == search {
      return true
    }
  }

  return false
}
