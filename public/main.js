/**
 * Add a date to the textarea
 */
function addDate() {
  const dateTimesTextArea = document.querySelector("#datetimes-textarea");
  const dateTimeSelect = document.querySelector("#datetime-select");
  const value = new Date(dateTimeSelect.value);
  dateTimeSelect.value = "";

  if (dateTimesTextArea.value === "") {
    dateTimesTextArea.value += value.toLocaleString();
  } else {
    dateTimesTextArea.value += "\n" + value.toLocaleString();
  }
}
