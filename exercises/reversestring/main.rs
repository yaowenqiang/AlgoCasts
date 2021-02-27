fn main() {

    println!("hello world");
    println!("{}", reverse("abc你好"))
}

fn reverse(s :  &str) -> &'static str {
    let mut retunstr = String.from;
    for i in s.chars() {
	retunstr = format!("{}{}",i, retunstr);
    }
    println!("{}", &retunstr);
    return  ""
}
