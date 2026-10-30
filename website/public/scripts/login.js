const loginBtn = document.getElementById("submit-login");

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
            console.error('Server error:', res.status);
        }
    } catch (error) {
        console.error('Network error:', error);
    }
}