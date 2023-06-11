import text from './text';

const input = {
	variants: {
		'primary-1': {
			field: {
				border: 'solid 1px',
				borderColor: 'gray.500',
				_focus: {
					border: 'none',
					outline: 'solid 3px',
					outlineColor: 'gray.600',
					outlineOffset: '-2px',
				},
				borderRadius: 'sm',
				...text.variants.h6,
				padding: '8px 16px',
			},
		},
	},
};

export default input;
