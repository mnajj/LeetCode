func countTime(time string) int {
  a, b, c, d: = 1, 1, 1, 1
  
  if time[4] == '?' {
    d = 10
  }
  
  if time[3] == '?' {
    c = 6
  }
  
  if time[1] == '?' {
    if time[0] == '2' {
      b = 4
    } else if time[0] == '?' {
      b = 1
    } else {
      b = 10
    }
  }
  
  if time[0] == '?' {
    if time[1] > '3' && time[1] <= '9' {
      a = 2
    } else if time[1] == '?' {
      a = 24
    } else {
      a = 3
    }
  }
  
  return a * b * c * d
}
