'use script';

window.onload = () => {
  const form = document.querySelector('form');
  const warning = document.querySelector('#warning');

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

    const reqBody = { username: username, password: password };


    // Scenario 2
    if (!username || !password) {
      warning.textContent = `All the input fields are required.`;
      // Scenario 4
    } else {

      fetch('/login', {
        method: 'POST',
        headers: { 'content-type': 'application/json' },
        body: JSON.stringify(reqBody),
      })
        .then(response => {
          if (response.status !== 200 && response.status !== 401) { throw new Error('no response'); }
          return response.json();
        })
        .then(message => {
          localStorage.setItem('token', message.token);
          console.log(`You're logged in.`);
          warning.textContent = `You're logged in.`;
          setTimeout(function () { window.location.replace('./index.html'); }, 1000);
        })
        .catch((err) => { console.log('error happend'); });
    }
  });
}