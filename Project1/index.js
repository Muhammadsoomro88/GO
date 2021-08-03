function validateForm(){
    let fname = document.getElementById("firster").value;
    let lname = document.getElementById("laster").value;

    if (fname == ""){
        document.getElementById("demo").innerHTML="Name field is required!";
        return false
    }
    if (lname == ""){
        document.getElementById("demo").innerHTML="Password field is required!";
        return false
    }
}
