package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
    "log"
)


type Sudoku struct {
    Grid[9][9]int
}


func main(){
  var base Sudoku
  var final = base.ReadFile()
  for i, _ := range final{
    final[i].Display()
    final[i].IsValid(0)
    fmt.Println()
  }

}

func (this Sudoku) ReadFile() []Sudoku {
  const path = "./grid/";
  var arrayReturn []Sudoku

  files, err := ioutil.ReadDir(path)
  if err != nil {
      log.Fatal(err)
  }

  for _, f := range files{
          b, err := ioutil.ReadFile(path + f.Name()) // just pass the file name
          if err != nil {
              fmt.Print(err)
          }

          str := string(b) // convert content to a 'string'
          var valuesInFiles = strings.Split(str, "\n")

          for index, _ := range valuesInFiles {
              var valuesInArray = strings.Split(valuesInFiles[index], "")
              for o, _ := range valuesInArray {
                  var value = 0
                  if(valuesInArray[o] == ".") {
                    value = 0
                  } else {
                    v, _ := strconv.Atoi(valuesInArray[o])
                    value = v
                  }
                  this.Grid[index][o] = value
              }

          }
          arrayReturn = append(arrayReturn, this)
  }
  return arrayReturn

}



func (this Sudoku) Display(){
  for kX, _ := range this.Grid {
        fmt.Println("")
        if (kX == 3 || kX == 6) {
            fmt.Println("-----------")
        }
        for kY, _ := range this.Grid[kX] {
            if(kY == 3 || kY == 6) {
                fmt.Print("|")
            }
            fmt.Print(this.Grid[kX][kY])
        }
    }
    fmt.Println("")
}

func (this Sudoku) NotOnLine(k int, x int) bool {
  var y int
    for y=0; y < 9; y++ {
        if (this.Grid[x][y] == k){
            return false
          }
        }
    return true
}

func (this Sudoku) NotOnRow(k int, y int) bool {
  var x int
  for x=0; x < 9; x++{
       if (this.Grid[x][y] == k){
           return false;
         }
       }
   return true;
}

func (this Sudoku) NotOnBloc(k int, x int, y int) bool {
  var firstX, firstY int;
  firstX =  x-(x%3)
  firstY =  y-(y%3)// coordonées de chaque première case de bloc
  for x = firstX; x < firstX+3; x++ {
        for y = firstY; y < firstY+3; y++ {
            if (this.Grid[x][y] == k){
                return false;
              }
        }
      }
    return true;
}
/*si x = 0,  x%3 = 0,  x - (x%3) = 0
  si x = 1,  x%3 = 1,  x - (x%3) = 0
  si x = 2,  x%3 = 2,  x - (x%3) = 0
  si x = 3,  x%3 = 0,  x - (x%3) = 3
  si x = 4,  x%3 = 1,  x - (x%3) = 3
  si x = 5,  x%3 = 2,  x - (x%3) = 3
  si x = 6,  x%3 = 0,  x - (x%3) = 6
  si x = 7,  x%3 = 1,  x - (x%3) = 6
  si x = 8,  x%3 = 2,  x - (x%3) = 6
Pareil avec y
Coordonnées : 0-0, 0-3 0-6, 3-0, 3-3 ,3-6, 6-0, 6-3, 6-6*/

func (this Sudoku)IsValid(position int) bool {

  if (position == 9*9){
      this.Display()
        return true;
      }

      var x, y, k int

    x = position/9 // Coordonnées de la case à position n exemple : position 51 x = 5
    y = position%9 // Coordonnées de la case à position n exemple : position 51 y = 6

    if (this.Grid[x][y] != 0){
        return this.IsValid(position+1);
      }

    for k=1; k <= 9; k++ { // test valeurs de 1 à 9
        if (this.NotOnLine(k,x) && this.NotOnRow(k,y) && this.NotOnBloc(k,x,y)){
            this.Grid[x][y] = k; // si tout est bon rajoute la valeur k dans la case

            if (this.IsValid(position+1)){//
                return true;
              }
        }
    }
    this.Grid[x][y] = 0;// si aucune valeur n'est bonne on remet la case à zéro
    return false;

}
