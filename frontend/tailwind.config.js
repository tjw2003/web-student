/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./src/**/*.{html,js,svelte,ts}",
    "./node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}",
  ],
  theme: {
    extend: {
      backgroundImage: {
        "tjw-sb": "url('/image/bg.jpg')",
      },
    },
  },
  plugins: [require("flowbite/plugin")],
  darkMode: 'class',
};
