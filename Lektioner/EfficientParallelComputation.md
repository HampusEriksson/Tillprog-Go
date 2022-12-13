https://yourbasic.org/golang/efficient-parallel-computation/

- Divide the work into units that take about 100μs to 1ms to compute.
  _ If the work units are too small, the adminis­trative over­head of divi­ding the problem and sched­uling sub-problems might be too large.
  _ If the units are too big, the whole computation may have to wait for a single slow work item to finish. This slowdown can happen for many reasons, such as scheduling, interrupts from other processes, and unfortunate memory layout.
  Note that the number of work units is independent of the number of CPUs.
- Try to minimize the amount of data sharing.
  - Concurrent writes can be very costly, particularly so if goroutines execute on separate CPUs.
  - Sharing data for reading is often much less of a problem.
- Strive for good locality when accessing data.
  - If data can be kept in cache memory, data loading and storing will be dramatically faster.
  - Once again, this is particularly important for writing.
