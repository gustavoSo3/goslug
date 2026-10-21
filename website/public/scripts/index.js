const newRedirectForm = document.getElementById("new-redirect-form");
const subBtn = document.getElementById("submit-new-redirect");

const redirectTable = document.getElementById("redirects-table");

newRedirectForm.addEventListener("submit", (event) => {
    event.preventDefault();

    var newRow = `
        <tr>
            <td>{{.Description}}</td>
            <td><a href="{{.Slug}}"><span style="color: gray;">{{$.Host}}</span>{{.Slug}}</a></td>
            <td><a href="{{.Destination}}">{{.Destination}}</a></td>
        </tr>
        `

    redirectTable.insertAdjacentHTML("beforeend", newRow);

})

