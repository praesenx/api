import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import tailwindcss from '@tailwindcss/vite'; // Import the Tailwind Vite plugin

// https://vitejs.dev/config/
export default defineConfig({
	define: {
		'process.env.VITE_API_URL': JSON.stringify(process.env.VITE_API_URL),
	},
	plugins: [vue(), tailwindcss()],
	resolve: {
		alias: [
			{
				find: /^~.+/,
				replacement: (val) => {
					return val.replace(/^~/, '');
				},
			},
		],
	},
});
