import { Subscription } from 'rxjs';
import { onMounted, onUnmounted } from 'vue'
import { type Signal } from '/@src/utils/enums'

type InternalEventHanlder = (data: unknown) => void;
type InternalEventListenersPair = [Signal, InternalEventHanlder]

export function useListener(event: Signal, handler: InternalEventHanlder) {
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
