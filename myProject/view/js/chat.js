
const chatMessages = document.getElementById('chatMessages');
const userInput = document.getElementById('userInput');
const sendButton = document.getElementById('sendButton');
var cookies = document.cookie;
console.log(cookies)
cookies = cookies.split(";");
console.log(cookies)
var cookie = cookies[0].trim();
console.log(cookie)
var cookieParts = cookie.split("="); 
var ming = cookieParts[2];
var value = cookieParts[1];
console.log("Name: " + ming + ", Value: " + value);
// Event listener for send button click
sendButton.addEventListener('click', sendMessage);

// Event listener for Enter key press
userInput.addEventListener('keypress', function(event) {
  if (event.key === 'Enter') {
    sendMessage();
  }
});

setTimeout(function() {
  appendBotReply(`Hello ${ming}! How can I assist you today?
  `);
}, 600);
// Call the function when the page loads

function sendMessage() {
  const question = userInput.value.trim();

  if (question === '') {
    return;
  }

  appendUserMessage( question);


  // Send the message to the backend
  fetch('/chat', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ question })
  })
    .then(function(response) {
      // if (response.ok) {
        return response.json();
      // } else {
        // throw new Error(response.statusText);

      // }
    })
    .then(function(data) {
      // Read the bot's reply from the response object
      const botReply = data.answer 
      const botReply1 = data.error
      console.log(botReply)
      if (typeof botReply === 'undefined') {
        var botReply2 = botReply1
      }else{
        var botReply2 = botReply
      }
  
      // console.log(data.error)
      setTimeout(function() {
        appendBotReply(botReply2);
      }, 600);
    
    })
    .catch(function(error) {
      console.error(error);
    });

  // Clear the input field
  userInput.value = '';
}

// Function to append a user message to the chat messages div
function appendUserMessage(message) {
  const newUserContainer = document.createElement('div');
  newUserContainer.classList.add('message-container');
  newUserContainer.classList.add('user-container');

  const userAvatar = document.createElement('img');
  if (value == 1){
    userAvatar.src = './images/profile/p1.png';
  }else if (value == 2){
    userAvatar.src = './images/profile/p2.png';
  }else if (value == 3){
    userAvatar.src = './images/profile/p3.png';
  }else if(value == 4){
    userAvatar.src = './images/profile/p4.png';
  }else{
    userAvatar.src = './images/bot.svg';
  }
 
  userAvatar.classList.add('user-avatar');

  const newUserMessage = document.createElement('div');
  newUserMessage.textContent = message;
  newUserMessage.classList.add('user-message');

  newUserContainer.appendChild(userAvatar);
  newUserContainer.appendChild(newUserMessage);
  chatMessages.appendChild(newUserContainer);
  chatMessages.scrollTop = chatMessages.scrollHeight;
}


// Function to append a bot reply to the chat messages div
function appendBotReply(message) {
  const newBotContainer = document.createElement('div');
  newBotContainer.classList.add('message-container');
  newBotContainer.classList.add('bot-container');

  const botAvatar = document.createElement('img');
  botAvatar.src = './images/bot.svg';
  botAvatar.classList.add('bot-avatar');

  const newBotMessage = document.createElement('div');
  newBotMessage.classList.add('bot-message');
  newBotMessage.textContent = message;

  newBotContainer.appendChild(botAvatar);
  newBotContainer.appendChild(newBotMessage);
  chatMessages.appendChild(newBotContainer);
  chatMessages.scrollTop = chatMessages.scrollHeight;
}

function Logout(){
  fetch("/logout")
  .then(res => {
      if (res.ok){
          window.open("index.html","_self")
      }else{
          throw new Error(res.statusText)
      }
  }).catch(e =>{
      alert(e)
  })
}