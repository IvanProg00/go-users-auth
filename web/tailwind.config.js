const colors = require('tailwindcss/colors');

module.exports = {
  purge: ['./src/**/*.{js,jsx,ts,tsx}', './public/index.html'],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {},
    colors: {
      primary: '#111',
      gray: colors.gray,
      button: '#fff'
    },
    backgroundColor: {
      gray: colors.gray,
      button: '#0D47A1',
      input: 'rgb(212, 212, 216, 0.25)'
    },
    borderRadius: {
      input: '5px'
    },
    borderColor: {
      input: 'rgba(59, 130, 246, 0.5)',
    },
    ringColor: {
      input: 'rgba(59, 130, 246, 0.5)',
      button: 'rgba(59, 130, 246, 0.5)'
    }
  },
  variants: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms')({
      strategy: 'class'
    }),
  ],
}
