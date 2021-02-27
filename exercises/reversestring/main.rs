fn main() {
    let s = String::from("hello world");
    println!("{}", reverse(s));
    let s = String::from("你好");
    println!("{}", reverse(s));
}

/*
fn reverse(s :  &str) -> &'static str {
    let mut retunstr = String.from(s);
    for i in s.chars() {
	retunstr = format!("{}{}",i, retunstr);
    }
    println!("{}", &retunstr);
    return  ""
}
*/
fn reverse(mut s: String) -> String {
    let mut ss = String::from("");
    while s.len() > 0 {
        //let c = s.pop().unwrap();
        ss.push(s.pop().unwrap());
    }
    return ss
}
