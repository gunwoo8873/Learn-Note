pub fn main() {
  if_controlflow();
  loop_controlflow();
  label_loop_controlflow();
  while_controlflow();
  for_controlflow();
}

fn if_controlflow() {
  let a = 15;
  let b = 20;

  if a > b {
    println!("{} > {}", a, b);
  }
  else {
    println!("{} < {}", a, b);
  }
}

fn loop_controlflow() {
  let mut default_count = 0;
  let max_count = 15;

  loop {
    println!("Loop Count : {}", default_count);

    if default_count == max_count {
      break;
    }

    default_count += 1;
  }
}

fn label_loop_controlflow() {
    let mut default_count = 0;

    //// Label Loop
    'counting_up: loop { // 'name : loop {...} = loop label?
        println!("count = {}", default_count);

        let mut remain = 10;

        loop {
            println!("remain = {}", remain);

            if remain == 9 {
                break;
            }

            if default_count == 5 {
                break 'counting_up; // Break after Return Label loop run
            }

            remain -= 1;
        }

        default_count += 1;
    }

    println!("End Count = {}", default_count);
}
// Note : 다중 Loop 에서 remain 이 10에서 9로 반환값을 출력을 할 때 default_count 가 0에서 5에 도달 할 때 까지 반복을 한다.
//        하지만 '의 사용처는 많아 단순히 특정 Loop 의 이용함이 아니라 문자열이나 타입 명시 할 경우에 사용하는 것 같다.

fn while_controlflow() {
    let mut default_count = 0;

    while default_count < 5 {
        default_count += 1;

        if default_count == 5 {
            break
        }
        println!("While End Result : {}", default_count);
    }

    let arr = ["A", "B", "C", "D", "E"];

    while default_count < 5 {
        println!("Array While End Result = {}", arr[default_count]); // arr[index] = name[index 0..5]
        default_count += 1;
    }
}

fn for_controlflow() {
    let arr = [10, 20, 30, 40, 50];

    for element in arr { // in arr = arr[index] bind
        println!("{}", element);
    }

    for element in (1..5).rev() { // in (1..5) = 1 ~ 4
        println!("{}", element);
    }
}
// Note : for은 컬렌션의 아이템에 대하여 임의의 코드를 수행시킬 수 있다.
//        rev는 Array의 Index 순서를 반대로 배치 후 출력된다.