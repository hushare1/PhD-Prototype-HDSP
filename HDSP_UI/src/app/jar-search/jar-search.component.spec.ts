import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { JarSearchComponent } from './jar-search.component';

describe('JarSearchComponent', () => {
  let component: JarSearchComponent;
  let fixture: ComponentFixture<JarSearchComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ JarSearchComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(JarSearchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
