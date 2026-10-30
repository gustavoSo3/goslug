const loginBtn = document.getElementById("submit-login");
const liveAlertDiv = document.getElementById('login-live-alert')

const usernameInput = document.getElementById("user-name");
const passwordInput = document.getElementById("password");

const loginURL = "/login";

loginBtn.addEventListener("click", async () => {
    if (AreValidFields(usernameInput.value, passwordInput.value)){   
        await postLogin(usernameInput.value, passwordInput.value);
    }
});

document.addEventListener("keydown", async (event) => {
    if (AreValidFields(usernameInput.value, passwordInput.value)){
        if(event.key === "Enter"){
            await postLogin(usernameInput.value, passwordInput.value);
        }
    } 
});


document.addEventListener("keyup", () => {
    if (AreValidFields(usernameInput.value, passwordInput.value)){
        loginBtn.disabled = false;
    } else {
        loginBtn.disabled = true;
    }
});


function AreValidFields(username, password) {

    username = username.trim();
    password = password.trim();

    // If any empty string
    if (username === "" || password === "" ) return false;


	// both user and password need to be less than 30 chars
    if (username.length > 30) return false;
    if (password.length > 30) return false;


	return true;
}

async function postLogin(username, password) {

    username = username.trim();
    password = password.trim();
    

    const login = {
        username: username,
        password: password,
    };

    username.value = "";
    password.value = "";

    try {
        // Make the POST request to /login
        const res = await fetch(loginURL, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(login) 
        });

        if (res.ok) {
            location.href = "/admin"
        } else {

            switch (res.status) {
                case 401:
                    appendAlert("Bad login. Double check your credentials.", "secondary");
                    break;
                default:
                    appendAlert("Server error: " + res.status, res.body, "warning");
                    break;
            }

        }
    } catch (error) {
        appendAlert("Server error: " + error , "danger");
    }
}


//login-live-alert
function appendAlert(message, type)  {
  const wrapper = document.createElement('div')
  wrapper.innerHTML = [
    `<div class="alert theme-${type}" role="alert">`,
    `   <p>${message}</p>`,
    '   <button type="button" class="btn-close ms-auto" data-bs-dismiss="alert" aria-label="Close"></button>',
    '</div>'
  ].join('')

  liveAlertDiv.append(wrapper)
}