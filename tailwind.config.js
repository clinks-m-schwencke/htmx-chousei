/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.templ", "./public/css/main.css"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
};
