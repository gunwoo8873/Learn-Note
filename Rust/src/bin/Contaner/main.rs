fn static_array() {
  let nomal_type_a = [1, 2, 3, 4, 5];
  println!("Nomal Array A = {:?}", nomal_type_a);

  let nomal_type_b : [i8; 5] = [1, 2, 3, 4, 5];
  println!("Nomal Array B = {:?}", nomal_type_b);

  let nomal_type_c  = [3; 5]; // Index value to all 3 create
  println!("Nomal Array C = {:?}", nomal_type_c);
}

fn main() {
  static_array()
}