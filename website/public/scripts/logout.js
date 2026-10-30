const logoutBtn = document.getElementById("logout-button");

const loginURL = "/login";

logoutBtn.addEventListener("click", async () => {

    // Make the DELETE request to /login
    const res = await fetch(loginURL, {
    method: 'DELETE',
    });

    if (res.ok) {
        location.href = "/"
    } else {
        console.error('Server error:', res.status);
        location.reload()
    }
});