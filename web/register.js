'use script';

window.onload = () => {
    const form = document.querySelector('form');
    const warning = document.querySelector('#warning');
    const firstButton = document.querySelector('#first')
    const lastField = form.lastElementChild;
    const secondButton = document.querySelector('#second');
    const newButton = document.createElement('a');
    newButton.innerHTML = '<button type="button">Proceed to the login page</button>';
    newButton.href = './login.html';
    newButton.id = 'first';

    //password visibility "button"
    const eyePic = document.querySelector('img[class="eye"]')
    const inputPassword = document.querySelector('input[name="password"]')
    let count = 0;
    eyePic.onclick = () => {
        if (count === 1) {
            eyePic.src = "./iconfinder_ic_visibility_48px_352209.svg"
            inputPassword.type = 'password';
            count = 0;
        } else {
            eyePic.src = '/iconfinder_ic_visibility_off_48px_3669412.svg'
            inputPassword.type = 'text';
            count = 1
        }
    }

    form.addEventListener('submit', (event) => {
        event.preventDefault();
        warning.textContent = '';

        let username = form.username.value;
        let password = form.password.value;
        let email = form.email.value;

        const reqBody = { username: username, password: password, email: email };

        // Scenario 2
        if (!username || !password || !email) {
            warning.textContent = `Username and password and email are required.`;
            // Scenario 4
        } else if (password.length < 8) {
            warning.textContent = `The password must be minimum 8 characters long.`;
        } else {

            fetch('/register', {
                method: 'POST',
                headers: { 'content-type': 'application/json' },
                body: JSON.stringify(reqBody),
            })
                .then(response => {
                    if (response.status !== 200 && response.status !== 400) { throw new Error('no response'); }
                    return response.json();
                })
                .then(message => {
                    console.log(message);
                    if (message.error) {
                        warning.textContent = message.error;
                        console.log(warnning.textContent);
                    } else {
                        console.log(`I'm in`);
                        // Scenario 5
                        warning.textContent = `Dear, ${message.username}, your registration is approved.\n Please validate it with the link sent on your given email address.`
                        lastField.removeChild(firstButton);
                        lastField.removeChild(secondButton);
                        lastField.appendChild(newButton);
                        lastField.appendChild(secondButton);

                    }

                })
                .catch((err) => { console.log('error happend'); });

        }

    });

}