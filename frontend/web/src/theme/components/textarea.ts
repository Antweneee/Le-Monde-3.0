import text from './text';

const textarea = {
	variants: {
		'primary-1': {
			border: 'solid 1px',
			borderColor: 'gray.500',
			_focus: {
				border: 'none',
				outline: 'solid 3px',
				outlineColor: 'gray.600',
				outlineOffset: '-2px',
			},
			borderRadius: 'sm',
			...text.variants.p,
			padding: '10px 16px',
		},
	},
};

export default textarea;
