import { Layout } from './Layout';
import { DataDisplay } from './Data';
import { Widget } from './Widget';

export interface DisplayType {
  id: number;
  createdAt: Date;
  updatedAt: Date;
  name: string;
  key: string;
  description: string;
  BoxId: string;
  layoutId: number;
  widgetId: number;
  displayData: DataDisplay[];
  widget: Widget;
  layout: Layout;
}
