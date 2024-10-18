function reveal(value) {
        let myDiv = document.getElementsByTagName("form");
        let buttons = document.getElementById("section").getElementsByTagName("button");
        for (let i = 0; i < myDiv.length; i++) {
                if (myDiv[i].id == value+"Form"){
                        buttons[i].className = "unblock"
                        myDiv[i].style.display = "flex"
                } else {
                        buttons[i].className = "block"
                        myDiv[i].style.display = "none"
                }
                console.log(myDiv[i].style)
        }
}