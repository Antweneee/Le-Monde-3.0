import text from './text';

const button = {
	variants: {
		'primary-1': {
			bg: 'primary.1',
			...text.variants.h4,
			fontWeight: 'semi-bold',
			border: 'solid 1px black',
			borderRadius: 'lg',
			padding: '32px 40px',
			width: '100%',
		},
		'primary-2': {
			bg: 'primary.2',
			...text.variants.h4,
			border: 'solid 1px black',
			borderRadius: 'lg',
			padding: '32px 40px',
			width: '100%',
		},
		'primary-3': {
			bg: 'primary.3',
			...text.variants.h4,
			border: 'solid 1px black',
			borderRadius: 'lg',
			padding: '32px 40px',
			width: '100%',
		},
		'primary-4': {
			bg: 'primary.4',
			...text.variants.h4,
			border: 'solid 1px black',
			borderRadius: 'lg',
			padding: '32px 40px',
			width: '100%',
		},
		'secondary-1': {
			bg: 'secondary.1',
			...text.variants.h6,
			color: 'white',
			borderRadius: 'sm',
			padding: '8px 16px',
			width: '100%',
		},
		'secondary-2': {
			bg: 'secondary.2',
			...text.variants.h6,
			color: 'white',
			borderRadius: 'sm',
			padding: '8px 16px',
			width: '100%',
		},
		'secondary-3': {
			bg: 'secondary.3',
			...text.variants.h6,
			color: 'white',
			borderRadius: 'sm',
			padding: '8px 16px',
			width: '100%',
		},
		'secondary-4': {
			bg: 'secondary.4',
			...text.variants.h6,
			borderRadius: 'sm',
			padding: '8px 16px',
			width: '100%',
		},
	},
};

export default button;
