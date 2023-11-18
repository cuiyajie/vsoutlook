import { filter, Subject, type Subscription, tap } from 'rxjs';

export type SignalBody<EventEnum> = {
  type: EventEnum,
  data: unknown
}

export class EventBus<EventEnum> {
  private events: Subject<SignalBody<EventEnum>>;
  constructor() {
    this.events = new Subject();
  }

  listenFull(matcher: (value: SignalBody<EventEnum>) => boolean, handler: (item: SignalBody<EventEnum>) => void): Subscription {
    return this.events
      .asObservable()
      .pipe(filter(matcher), tap(handler))
      .subscribe();
  }

  listen(type: EventEnum, handler: (data: any) => void): Subscription {
    return this.listenFull(
      t => t.type === type,
      event => handler(event.data)
    )
  }

  trigger(type: EventEnum, data: any = null) {
    this.events.next({ type, data });
  }
}

export const bus = new EventBus();
