window = function () {
    fetch('/')
        .then(response => response.text())
}

function adduser() {
    console.log('bug');
    //reading the form data
    var data = {
        username: document.getElementById("name").value,
        email: document.getElementById("email").value,
        password: document.getElementById("password").value,
        avatar : document.querySelector('input[name="avatar"]:checked').value
    }
    console.log(data);
    var name = data.username
    if (name == '') {
        alert("enter your name")
        return
    } else if (data.email == '') {
        alert('Sorry email field is required')
        return
    } else if (data.password == '') {
        alert('password is highly recommendated for security purpose')
        return
    }
    fetch('/signup', {
        method: "POST",
        body: JSON.stringify(data),
        headers: { "Content-Type": "application/json; charset=UTF-8" }
    }).then(response => {
        if (response.status === 201) {
            alert("Successful");
            window.location.href = "login.html";
            // You can perform any additional actions or redirect to another page here
        } else {
            throw new Error(response.statusText);
        }
    }).catch(e => {
        alert(e);
    });
    
}


