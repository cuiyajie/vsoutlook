import { Subscription } from 'rxjs';
import { onMounted, onUnmounted } from 'vue'
import { type Signal } from '@src/utils/enums'

type InternalEventHanlder<T = any> = (data: T) => void;
type InternalEventListenersPair = [Signal, InternalEventHanlder]

export function useListener<T = any>(event: Signal, handler: InternalEventHanlder<T>) {
  let listener: Subscription | null = null;
  onMounted(() => listener = bus.listen(event, handler))
  onUnmounted(() => listener?.unsubscribe())
}

export function useListeners(rows: InternalEventListenersPair[]) {
  const listeners: Subscription = new Subscription();
  onMounted(() => {
    rows.forEach(([event, handler]) => {
      listeners.add(bus.listen(event, handler))
    })
  })
  onUnmounted(() => listeners.unsubscribe())
}
