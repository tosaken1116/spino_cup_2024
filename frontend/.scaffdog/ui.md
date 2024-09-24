---
name: "ui"
root: "src/components/ui"
output: "**/*"
ignore: []
questions:
  name: "Please enter component name"
---

# Variables

- componentName :`{{ inputs.name | pascal }}`

# `{{ componentName }}/index.tsx`

```tsx
import type { FC } from 'react';
type Props = {
  children: React.ReactNode;
};

export const {{ componentName }}: FC<Props> = ({ children }) => {
  return <div>{children}</div>;
};
```

# `{{ componentName }}/index.stories.tsx`

```tsx
import { Meta } from '@storybook/react';
import { {{ componentName }} } from '.';

export default {
  title: 'ui/{{ componentName }}',
  component: {{ componentName }},
} satisfies Meta<typeof {{ componentName }}>

export const Default = {};
```
