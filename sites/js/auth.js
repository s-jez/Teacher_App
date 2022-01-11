const formRegisterUser = document.getElementById('formRegisterUser');
const formLoginUser = document.getElementById('formLoginUser');

// POST
formRegisterUser.addEventListener('submit', async (event) => {
    event.preventDefault();
    let userData = {
        UserName: document.getElementById("fusername").value,
        Password: document.getElementById("fpassword").value,
        Email: document.getElementById('femail').value,
        RoleID: document.getElementById('frole').value
    }
    let userRole = parseInt(userData.RoleID);
    if (!userData.UserName || !userData.Password || !userData.Email || userRole == 0) { alert('Please enter data to form'); return; }
    const validateEmail = (email) => {
        return String(email)
            .toLowerCase()
            .match(
                /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
            );
    }
    if (!validateEmail(userData.Email)) {
        alert('Email incorrect!');
        return
    }
    fetch(`/register`, {
        method: "POST",
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "UserName": userData.UserName,
            "Password": userData.Password,
            "Email": userData.Email,
            "RoleID": userRole
        })
    }).then(async (result) => {
        const content = await result.json();
        if (result) {
            document.getElementById("fusername").value = "";
            document.getElementById("fpassword").value = "";
            document.getElementById('femail').value = "";
            document.getElementById('frole').value = null;
            // hide modal window
            let modalWindow = document.querySelector('#exampleModal4');
            modalWindow.style.display = 'none';
            let divBackground = document.querySelector('.modal-backdrop');
            divBackground.remove();

            alert('Created user successfully! \nNow login to your account!');
            console.log(content);
        } else {
            alert('Error deleting user');
            return
        }
    }).catch(e => {
        console.log(e);
    });
})
// POST
formLoginUser.addEventListener('submit', async (event) => {
    event.preventDefault();
    let userData = {
        UserName: document.getElementById("username").value,
        Password: document.getElementById("password").value,
        Email: document.getElementById('email').value
    }
    if (!userData.UserName || !userData.Password || !userData.Email) { alert('Please enter data to form'); return; }
    fetch(`/login`, {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('access_token')
        },
        body: JSON.stringify(userData)
    }).then(async (result) => {
        const content = await result.json();
        if (result) {
            document.getElementById("username").value = "";
            document.getElementById("password").value = "";
            document.getElementById('email').value = "";
            // hide modal window
            let modalWindow = document.querySelector('#exampleModal5');
            modalWindow.style.display = 'none';
            let divBackground = document.querySelector('.modal-backdrop');
            divBackground.remove();
            // alert
            alert('User logged in successfully! \nWelcome to Student App');
            location.href = '/logged';
            console.log(content.AccessToken);
            localStorage.setItem("access_token", JSON.stringify(content.AccessToken.Token).slice(1, -1));
            localStorage.setItem("refresh_token", JSON.stringify(content.RefreshToken.Token).slice(1, -1));
        }
    }).catch((err) => {
        console.log(err)
    })
})