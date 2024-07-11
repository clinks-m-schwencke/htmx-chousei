/**
 * Add a date to the textarea
 */
function addDate() {
  const dateTimesTextArea = document.querySelector("#datetimes-textarea");
  const datePicker = document.querySelector("#date-picker");
  const timePicker = document.querySelector("#time-picker");
  console.log(datePicker.value)
  console.log(timePicker.value)
  const value = new Date(datePicker.value + "T" + timePicker.value);

  let localeString = ''
  if(value.toLocaleString() === 'Invalid Date') {
    // Attempt with only date
    const value = new Date(datePicker.value)
    if(value.toLocaleString() !== 'Invalid Date'){
      localeString = value.toLocaleDateString()
    }else {
      return
    }
  } else {
    localeString = value.toLocaleString()
  }

  if (dateTimesTextArea.value === "") {
    dateTimesTextArea.value += localeString;
  } else {
    dateTimesTextArea.value += "\n" + localeString;
  }
}
