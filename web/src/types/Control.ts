import { DataControl } from './Data';
import { Layout } from './Layout';
import { Widget } from './Widget';

export interface Control {
  id: number;
  createdAt: Date;
  updatedAt: Date;
  name: string;
  key: string;
  description: string;
  BoxId: string;
  layoutId: number;
  widgetId: number;
  controlData: DataControl[];
  widget: Widget;
  layout: Layout;
}
