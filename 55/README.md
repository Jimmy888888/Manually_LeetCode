greedy : in every step, choose the biggest number

success condition : the position reach the last index

## values :

- step : the step of the current position

- newStep : the step of the next position

- position : the position of the current position


## how to find the next position

#### way 1 : simlply find the biggest number in the range of step

```GO
for j := i+1; j <= i + step; j++ {
    if newStep < nums[j] {
        newStep = nums[j]
        position = j
    }
}
```

result : fail, because the partly biggest number may not be the longest step

#### way 2 : include the distance, using the (i + step) as the stander

```GO
for j := i+1; j <= i + step; j++ {
    if j + nums[j] >= i + step {
        newStep = nums[j]
        position = j
        break
    }
}
```
result : success


