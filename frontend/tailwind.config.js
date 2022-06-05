
// See default config https://github.com/tailwindcss/tailwindcss/blob/master/stubs/defaultConfig.stub.js
module.exports = {
  theme: {
    extend: {
		fontFamily: {
			sans: ["Inter var", "-apple-system", "BlinkMacSystemFont", "Segoe UI", "Roboto", "Helvetica Neue", "Arial", "Noto Sans", "sans-serif", "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji"]
		},
		colors: {
			primary: '#FF6900',
			'primary-dark': '#CC5500',
		},
		borderRadius: {
			'card': '1.5rem',
		},
	}
  },
  variants: {}
}