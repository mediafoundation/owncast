import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import { ActionButtonRow } from './ActionButtonRow';
import { ActionButton } from '../ActionButton/ActionButton';

export default {
  title: 'owncast/Components/Action Buttons/Buttons Row',
  component: ActionButtonRow,
  parameters: {
    docs: {
      description: {
        component: `This is a horizontal row of buttons that could be statically created by the Owncast application (such as Notify, Follow) or are user-generated external actions (Donate, Learn more, etc).
        There can be any number of buttons, including zero. They should wrap if needed and handle resizing.`,
      },
    },
  },
} as ComponentMeta<typeof ActionButtonRow>;

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const Template: ComponentStory<typeof ActionButtonRow> = args => {
  const { buttons } = args as any;
  return <ActionButtonRow>{buttons}</ActionButtonRow>;
};

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const actions = [
  {
    url: 'https://owncast.online/docs',
    title: 'Documentation',
    description: 'Owncast Documentation',
    icon: 'https://owncast.online/images/logo.svg',
    color: '#5232c8',
    openExternally: false,
  },
  {
    url: 'https://opencollective.com/embed/owncast/donate',
    title: 'Support Owncast',
    description: 'Contribute to Owncast',
    icon: 'https://opencollective.com/static/images/opencollective-icon.svg',
    color: '#2b4863',
    openExternally: false,
  },
];

const itemSelected = a => {
  console.log('itemSelected', a);
  action(a.title);
};

const buttons = actions.map(a => <ActionButton externalActionSelected={itemSelected} action={a} />);
export const Example1 = Template.bind({});
Example1.args = {
  buttons,
};
